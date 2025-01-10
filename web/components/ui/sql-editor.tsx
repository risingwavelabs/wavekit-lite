import { useState, useCallback, useEffect, useRef } from 'react'
import Editor, { useMonaco } from '@monaco-editor/react'
import type { languages } from 'monaco-editor'
import { Button } from "@/components/ui/button"
import { Play, X, Plus } from 'lucide-react'
import { GenerateQuery } from "@/components/ui/generate-query"
import { DatabaseInsight } from "@/components/ui/database-insight"
import { RisingWaveNodeData } from "@/components/streaming-graph"
import { useTheme } from 'next-themes'

// Move these to a separate constants file if needed
const SQL_COMPLETIONS = {
  keywords: ['SELECT', 'FROM', 'WHERE', 'GROUP BY', 'ORDER BY', 'HAVING', 'JOIN', 'LEFT JOIN', 'RIGHT JOIN', 'INNER JOIN', 'LIMIT', 'OFFSET'],
  functions: ['COUNT', 'SUM', 'AVG', 'MAX', 'MIN', 'COALESCE', 'NULLIF', 'CAST', 'DATE', 'EXTRACT'],
  operators: ['AND', 'OR', 'NOT', 'IN', 'BETWEEN', 'LIKE', 'IS NULL', 'IS NOT NULL'],
}

interface EditorTab {
  id: string
  name: string
  content: string
  isDirty?: boolean
}

interface SQLEditorProps {
  width: number
  savedQueries: Array<{ id: string, name: string }>
  onRunQuery?: (query: string) => Promise<{ type: 'success' | 'error', message: string, rows?: Record<string, string>[] }>
  onSaveQuery?: (query: string, name: string) => void
  databaseSchema?: RisingWaveNodeData[]
  selectedDatabaseId?: string | null
  onCancelProgress?: (ddlId: string) => void
}

// Sample AI-generated queries
const SAMPLE_AI_QUERIES = [
  `-- Find active users with recent purchases
SELECT DISTINCT u.id, u.name, u.email
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE u.status = 'active'
  AND o.created_at >= DATE_SUB(NOW(), INTERVAL 1 MONTH)
ORDER BY o.created_at DESC;`,

  `-- Calculate monthly revenue by product
SELECT 
  p.name as product_name,
  DATE_FORMAT(o.created_at, '%Y-%m') as month,
  COUNT(*) as total_orders,
  SUM(o.quantity) as units_sold,
  SUM(o.quantity * p.price) as revenue
FROM orders o
JOIN products p ON o.product_id = p.id
GROUP BY p.name, DATE_FORMAT(o.created_at, '%Y-%m')
ORDER BY month DESC, revenue DESC;`,

  `-- Track user engagement metrics
SELECT 
  e.user_id,
  u.name,
  COUNT(DISTINCT DATE(e.event_date)) as active_days,
  COUNT(*) as total_events,
  STRING_AGG(DISTINCT e.event_type) as event_types
FROM events e
JOIN users u ON e.user_id = u.id
WHERE e.event_date >= DATE_SUB(NOW(), INTERVAL 7 DAY)
GROUP BY e.user_id, u.name
HAVING COUNT(*) > 5
ORDER BY active_days DESC;`
]

// Sample database schema for streaming graph visualization
const SAMPLE_SCHEMA: RisingWaveNodeData[] = [
  {
    id: 1,
    name: 'user_events',
    type: 'source',
    columns: [
      { name: 'user_id', type: 'INT', isPrimary: true },
      { name: 'event_type', type: 'VARCHAR' },
      { name: 'timestamp', type: 'TIMESTAMP' }
    ],
    connector: {
      type: 'kafka',
      properties: {
        topic: 'user_events',
        bootstrap_servers: 'kafka:9092'
      }
    }
  },
  {
    id: 2,
    name: 'user_metrics',
    type: 'materialized_view',
    columns: [
      { name: 'user_id', type: 'INT', isPrimary: true },
      { name: 'event_count', type: 'INT' },
      { name: 'last_seen', type: 'TIMESTAMP' }
    ],
    dependencies: [1]
  },
  {
    id: 3,
    name: 'product_events',
    type: 'source',
    columns: [
      { name: 'product_id', type: 'INT', isPrimary: true },
      { name: 'event_type', type: 'VARCHAR' },
      { name: 'timestamp', type: 'TIMESTAMP' }
    ],
    connector: {
      type: 'kafka',
      properties: {
        topic: 'product_events',
        bootstrap_servers: 'kafka:9092'
      }
    }
  },
  {
    id: 4,
    name: 'product_analytics',
    type: 'materialized_view',
    columns: [
      { name: 'product_id', type: 'INT', isPrimary: true },
      { name: 'view_count', type: 'INT' },
      { name: 'purchase_count', type: 'INT' }
    ],
    dependencies: [3]
  },
  {
    id: 5,
    name: 'user_product_recommendations',
    type: 'materialized_view',
    columns: [
      { name: 'user_id', type: 'INT' },
      { name: 'product_id', type: 'INT' },
      { name: 'score', type: 'FLOAT' }
    ],
    dependencies: [1, 3]
  }
]

export function SQLEditor({ width, savedQueries, onRunQuery, onSaveQuery, databaseSchema = SAMPLE_SCHEMA, selectedDatabaseId, onCancelProgress }: SQLEditorProps) {
  const { theme } = useTheme()
  const [mounted, setMounted] = useState(false)
  const [tabs, setTabs] = useState<EditorTab[]>(() => {
    if (typeof window !== 'undefined') {
      const savedTabs = localStorage.getItem('editor-tabs')
      if (savedTabs) {
        try {
          return JSON.parse(savedTabs)
        } catch (e) {
          console.error('Failed to parse saved tabs:', e)
        }
      }
    }
    return [{
      id: '1',
      name: 'Query 1',
      content: '',
      isDirty: false
    }]
  })
  const [activeTab, setActiveTab] = useState(() => {
    if (typeof window !== 'undefined') {
      const savedActiveTab = localStorage.getItem('editor-active-tab')
      if (savedActiveTab && tabs.some(tab => tab.id === savedActiveTab)) {
        return savedActiveTab
      }
    }
    return '1'
  })
  const [editingTabId, setEditingTabId] = useState<string | null>(null)
  const [editingName, setEditingName] = useState('')
  const [editorHeight, setEditorHeight] = useState('60%')
  const [graphHeight, setGraphHeight] = useState<string>('30vh')
  const [isResizingHeight, setIsResizingHeight] = useState(false)
  const [isGenerating, setIsGenerating] = useState(false)
  const [generateError, setGenerateError] = useState<string>("")
  const [queryResult, setQueryResult] = useState<{ type: 'success' | 'error', message: string, rows?: any[] }>()
  const [executionHistory, setExecutionHistory] = useState<Array<{
    query: string;
    timestamp: string;
    status: 'success' | 'error';
    message: string;
    rowsAffected?: number;
  }>>([])
  const [activeResultTab, setActiveResultTab] = useState<'result' | 'graph' | 'progress' | 'history'>('result')

  const editorRef = useRef<any>(null)
  const editInputRef = useRef<HTMLInputElement>(null)
  const monaco = useMonaco()

  const calculateGraphHeight = useCallback(() => {
    if (typeof editorHeight === 'string' && editorHeight.endsWith('%')) {
      const percentage = parseInt(editorHeight)
      // Account for essential heights only: toolbar(48px) + tab bar(41px) + generate query bar(56px)
      const otherElementsHeight = 48 + 41 + 56
      const remainingHeightVh = 100 - percentage
      const remainingHeightPx = (window.innerHeight * remainingHeightVh) / 100
      const actualGraphHeight = remainingHeightPx - otherElementsHeight
      setGraphHeight(`${actualGraphHeight}px`)
    }
  }, [editorHeight])

  // Calculate graph height based on editor height changes
  useEffect(() => {
    calculateGraphHeight()
  }, [calculateGraphHeight])

  // Add window resize listener to recalculate dimensions
  useEffect(() => {
    const handleResize = () => {
      calculateGraphHeight()

      // Trigger Monaco editor layout update
      if (editorRef.current) {
        editorRef.current.layout()
      }
    }

    window.addEventListener('resize', handleResize)
    return () => window.removeEventListener('resize', handleResize)
  }, [calculateGraphHeight])

  // Save active tab to localStorage
  useEffect(() => {
    localStorage.setItem('editor-active-tab', activeTab)
  }, [activeTab])

  // Remove the old useEffect for loading tabs since we're doing it in the initial state
  // Keep only the save effect
  useEffect(() => {
    localStorage.setItem('editor-tabs', JSON.stringify(tabs))
  }, [tabs])

  // Handle Ctrl+S
  useEffect(() => {
    const handleSave = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === 's') {
        e.preventDefault()
        const currentTab = tabs.find(tab => tab.id === activeTab)
        if (currentTab) {
          onSaveQuery?.(currentTab.content, currentTab.name)
          setTabs(prev => prev.map(tab =>
            tab.id === activeTab ? { ...tab, isDirty: false } : tab
          ))
        }
      }
    }

    window.addEventListener('keydown', handleSave)
    return () => window.removeEventListener('keydown', handleSave)
  }, [activeTab, tabs, onSaveQuery])

  // Configure Monaco Editor
  useEffect(() => {
    if (monaco) {
      const disposable = monaco.languages.registerCompletionItemProvider('sql', {
        provideCompletionItems: (model, position) => {
          const word = model.getWordUntilPosition(position)
          const range = {
            startLineNumber: position.lineNumber,
            endLineNumber: position.lineNumber,
            startColumn: word.startColumn,
            endColumn: word.endColumn,
          }

          const suggestions: languages.CompletionItem[] = []

          // Add suggestions for keywords, functions, operators
          SQL_COMPLETIONS.keywords.forEach(keyword => {
            if (keyword.startsWith(word.word.toUpperCase())) {
              suggestions.push({
                label: keyword,
                kind: monaco.languages.CompletionItemKind.Keyword,
                insertText: keyword,
                range,
                preselect: true,
              })
            }
          })

          // ... Add other suggestions (functions, operators) similarly

          return { suggestions }
        },
        triggerCharacters: [' ', '.', '('],
      })

      return () => disposable.dispose()
    }
  }, [monaco])

  // Tab management handlers
  const handleNewTab = () => {
    const newId = `${Date.now()}`
    setTabs(prev => [...prev, {
      id: newId,
      name: `Query ${prev.length + 1}`,
      content: '',
      isDirty: false
    }])
    setActiveTab(newId)
  }

  const handleCloseTab = (tabId: string, e: React.MouseEvent) => {
    e.stopPropagation()
    if (tabs.length === 1) return

    setTabs(prev => prev.filter(tab => tab.id !== tabId))
    if (activeTab === tabId) {
      const index = tabs.findIndex(tab => tab.id === tabId)
      const newActiveTab = tabs[index - 1]?.id || tabs[index + 1]?.id
      setActiveTab(newActiveTab)
    }
  }

  const handleEditorChange = (value: string | undefined) => {
    if (!value) return
    setTabs(prev => prev.map(tab =>
      tab.id === activeTab
        ? { ...tab, content: value, isDirty: true }
        : tab
    ))
  }

  // Tab drag and drop handlers
  const handleDragStart = (e: React.DragEvent, tab: EditorTab) => {
    e.dataTransfer.setData('text/plain', tab.id)
    const target = e.target as HTMLElement
    target.classList.add('opacity-50')
  }

  const handleDragEnd = (e: React.DragEvent) => {
    const target = e.target as HTMLElement
    target.classList.remove('opacity-50')
  }

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault()
    e.dataTransfer.dropEffect = 'move'
  }

  const handleDrop = (e: React.DragEvent, targetTab: EditorTab) => {
    e.preventDefault()
    const draggedTabId = e.dataTransfer.getData('text/plain')
    const draggedTab = tabs.find(tab => tab.id === draggedTabId)

    if (!draggedTab || draggedTab.id === targetTab.id) return

    const draggedIndex = tabs.findIndex(tab => tab.id === draggedTab.id)
    const targetIndex = tabs.findIndex(tab => tab.id === targetTab.id)

    setTabs(prev => {
      const newTabs = [...prev]
      newTabs.splice(draggedIndex, 1)
      newTabs.splice(targetIndex, 0, draggedTab)
      return newTabs
    })
  }

  // Tab name editing handlers
  const handleTabDoubleClick = (tab: EditorTab) => {
    setEditingTabId(tab.id)
    setEditingName(tab.name)
  }

  const handleTabNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEditingName(e.target.value)
  }

  const handleTabNameSave = () => {
    if (editingTabId && editingName.trim()) {
      setTabs(prev => prev.map(tab =>
        tab.id === editingTabId ? { ...tab, name: editingName.trim() } : tab
      ))
    }
    setEditingTabId(null)
  }

  const handleTabNameKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      handleTabNameSave()
    } else if (e.key === 'Escape') {
      setEditingTabId(null)
    }
  }

  // Focus input when editing starts
  useEffect(() => {
    if (editingTabId && editInputRef.current) {
      editInputRef.current.focus()
      // Use setTimeout to ensure selection happens after focus
      setTimeout(() => {
        editInputRef.current?.select()
      }, 0)
    }
  }, [editingTabId])

  // Handle vertical resizing
  const handleMouseDownVertical = useCallback((e: React.MouseEvent) => {
    e.preventDefault()
    const startY = e.clientY
    const startHeight = parseFloat(editorHeight)

    const handleMouseMove = (e: MouseEvent) => {
      const deltaY = e.clientY - startY
      const containerHeight = window.innerHeight
      const deltaPercent = (deltaY / containerHeight) * 100
      const newHeight = startHeight + deltaPercent

      if (newHeight >= 0 && newHeight <= 80) {
        setEditorHeight(`${newHeight}%`)
      }
    }

    const handleMouseUp = () => {
      document.removeEventListener('mousemove', handleMouseMove)
      document.removeEventListener('mouseup', handleMouseUp)
      document.body.style.cursor = 'default'
    }

    document.addEventListener('mousemove', handleMouseMove)
    document.addEventListener('mouseup', handleMouseUp)
    document.body.style.cursor = 'ns-resize'
  }, [editorHeight])

  // Remove the old handlers since we're handling everything in mousedown
  const handleMouseUpVertical = useCallback(() => {
    setIsResizingHeight(false)
    document.body.style.cursor = 'default'
  }, [])

  const handleMouseMoveVertical = useCallback((e: MouseEvent) => {
    if (!isResizingHeight) return

    const containerHeight = window.innerHeight
    const mouseY = e.clientY
    const percentage = (mouseY / containerHeight) * 100

    if (percentage >= 30 && percentage <= 80) {
      setEditorHeight(`${percentage}%`)
    }
  }, [isResizingHeight])

  useEffect(() => {
    if (isResizingHeight) {
      window.addEventListener('mousemove', handleMouseMoveVertical)
      window.addEventListener('mouseup', handleMouseUpVertical)
    }

    return () => {
      window.removeEventListener('mousemove', handleMouseMoveVertical)
      window.removeEventListener('mouseup', handleMouseUpVertical)
    }
  }, [isResizingHeight, handleMouseMoveVertical, handleMouseUpVertical])

  const handleRunQuery = useCallback(async () => {
    if (!editorRef.current) return

    // Get the selected text or full content
    const editor = editorRef.current
    const selection = editor.getSelection()
    const query = selection && !selection.isEmpty()
      ? editor.getModel().getValueInRange(selection)
      : tabs.find(tab => tab.id === activeTab)?.content

    if (!query?.trim()) return

    if (!selectedDatabaseId) {
      setQueryResult({
        type: 'error',
        message: 'Please select a database first'
      })
      return
    }

    try {
      const result = await onRunQuery?.(query)
      if (result) {
        setQueryResult(result)
        setExecutionHistory(prev => [{
          query,
          timestamp: new Date().toISOString(),
          status: result.type,
          message: result.message,
          rowsAffected: result.rows?.length
        }, ...prev])
        
        // Switch to appropriate tab based on result
        if (result.rows && result.rows.length > 0) {
          setActiveResultTab('result')
        } else {
          setActiveResultTab('history')
        }
      }
    } catch (error) {
      const errorResult = {
        type: 'error' as const,
        message: error instanceof Error ? error.message : 'Failed to execute query'
      }
      setQueryResult(errorResult)
      setExecutionHistory(prev => [{
        query,
        timestamp: new Date().toISOString(),
        status: 'error',
        message: errorResult.message
      }, ...prev])
      setActiveResultTab('history')
    }
  }, [activeTab, tabs, onRunQuery, selectedDatabaseId, editorRef])

  const handleGenerateQuery = useCallback((prompt: string) => {
    if (!prompt.trim() || !editorRef.current) return

    setIsGenerating(true)
    setGenerateError("")

    // Simulate AI query generation with a delay
    setTimeout(() => {
      try {
        const editor = editorRef.current
        const randomQuery = SAMPLE_AI_QUERIES[Math.floor(Math.random() * SAMPLE_AI_QUERIES.length)]
        const position = editor.getPosition()
        const lineContent = editor.getModel().getLineContent(position.lineNumber)
        const isEmptyLine = !lineContent.trim()

        // Add newlines if needed
        let queryToInsert = randomQuery
        if (!isEmptyLine) {
          queryToInsert = `\n\n${randomQuery}`
        }

        // Insert the query at cursor position
        editor.executeEdits('ai-generation', [{
          range: {
            startLineNumber: position.lineNumber,
            startColumn: position.column,
            endLineNumber: position.lineNumber,
            endColumn: position.column
          },
          text: queryToInsert,
        }])

        // Update tab content
        const newContent = editor.getValue()
        setTabs(prev => prev.map(tab =>
          tab.id === activeTab ? { ...tab, content: newContent, isDirty: true } : tab
        ))

        // Move cursor to the end of the inserted query
        const newPosition = editor.getPosition()
        editor.setPosition(newPosition)
        editor.focus()
      } catch (error) {
        setGenerateError(error instanceof Error ? error.message : "Failed to generate query")
      } finally {
        setIsGenerating(false)
      }
    }, 500)
  }, [activeTab])

  useEffect(() => {
    setMounted(true)
  }, [])

  return (
    <div className="flex flex-col h-full">
      <div className="flex items-center space-x-2 mb-2 p-2">
        <Button size="sm" variant="default" onClick={handleRunQuery}>
          <Play className="w-4 h-4 mr-1" />
          Run
        </Button>
      </div>

      <div className="border-b bg-muted/30 px-2">
        <div className="flex items-center">
          {tabs.map(tab => (
            <div
              key={tab.id}
              onClick={() => setActiveTab(tab.id)}
              draggable
              onDragStart={(e) => handleDragStart(e, tab)}
              onDragEnd={handleDragEnd}
              onDragOver={handleDragOver}
              onDrop={(e) => handleDrop(e, tab)}
              className={`group flex items-center gap-2 px-4 py-2 border-r cursor-pointer hover:bg-muted/50 ${activeTab === tab.id ? 'bg-background border-b-2 border-b-primary' : 'text-muted-foreground'
                }`}
            >
              <div className="flex items-center gap-1">
                {editingTabId === tab.id ? (
                  <input
                    ref={editInputRef}
                    value={editingName}
                    onChange={handleTabNameChange}
                    onBlur={handleTabNameSave}
                    onKeyDown={handleTabNameKeyDown}
                    className="bg-transparent border-none outline-none focus:ring-1 focus:ring-primary text-sm px-0 w-[96px] h-5 -my-[1px]"
                    onClick={e => e.stopPropagation()}
                  />
                ) : (
                  <span
                    className="text-sm select-none"
                    onDoubleClick={() => handleTabDoubleClick(tab)}
                  >
                    {tab.name}
                  </span>
                )}
                {tab.isDirty && (
                  <span className="h-1.5 w-1.5 rounded-full bg-primary/70" />
                )}
              </div>
              {tabs.length > 1 && (
                <button
                  onClick={(e) => handleCloseTab(tab.id, e)}
                  className="opacity-0 group-hover:opacity-100 hover:text-red-500"
                >
                  <X className="h-3 w-3" />
                </button>
              )}
            </div>
          ))}
          <button
            onClick={handleNewTab}
            className="p-2 hover:bg-muted/50"
          >
            <Plus className="h-4 w-4 text-muted-foreground" />
          </button>
        </div>
      </div>

      <div style={{ height: editorHeight }} className="relative">
        <Editor
          height="100%"
          defaultLanguage="sql"
          theme={mounted ? (theme === 'dark' ? 'vs-dark' : 'light') : 'light'}
          value={tabs.find(tab => tab.id === activeTab)?.content}
          onChange={handleEditorChange}
          onMount={(editor) => editorRef.current = editor}
          options={{
            minimap: { enabled: false },
            fontSize: 14,
            scrollBeyondLastLine: false,
            automaticLayout: true,
            padding: { top: 16 },
            quickSuggestions: {
              other: true,
              comments: true,
              strings: true
            },
            suggestOnTriggerCharacters: true,
            acceptSuggestionOnEnter: "on",
            tabCompletion: "on",
            suggest: {
              preview: true,
              showIcons: true,
              showStatusBar: true,
              showInlineDetails: true,
            }
          }}
          className="h-full"
          width={width}
        />
        <div
          className="absolute left-0 right-0 bottom-[-8px] h-[16px] z-10 cursor-ns-resize group"
          onMouseDown={handleMouseDownVertical}
        >
          <div className="absolute left-0 right-0 top-[7px] h-[2px] bg-border group-hover:bg-primary/20 group-active:bg-primary/40" />
        </div>
      </div>

      <div className="flex-1 min-h-0 flex flex-col">
        <GenerateQuery
          onGenerate={handleGenerateQuery}
          isGenerating={isGenerating}
          error={generateError}
          className="border-b bg-background/95"
        />
        <DatabaseInsight
          height={graphHeight}
          databaseSchema={databaseSchema}
          result={queryResult}
          selectedDatabaseId={selectedDatabaseId}
          onCancelProgress={onCancelProgress}
          executionHistory={executionHistory}
          activeTab={activeResultTab}
          onTabChange={setActiveResultTab}
        />
      </div>
    </div>
  )
} 