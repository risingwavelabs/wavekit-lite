from http import HTTPStatus
from typing import Any, Optional, Union

import httpx

from ... import errors
from ...client import AuthenticatedClient, Client
from ...models.test_database_connection_payload import TestDatabaseConnectionPayload
from ...models.test_database_connection_result import TestDatabaseConnectionResult
from ...types import Response


def _get_kwargs(
    *,
    body: TestDatabaseConnectionPayload,
) -> dict[str, Any]:
    headers: dict[str, Any] = {}

    _kwargs: dict[str, Any] = {
        "method": "post",
        "url": "/databases/test-connection",
    }

    _body = body.to_dict()

    _kwargs["json"] = _body
    headers["Content-Type"] = "application/json"

    _kwargs["headers"] = headers
    return _kwargs


def _parse_response(
    *, client: Union[AuthenticatedClient, Client], response: httpx.Response
) -> Optional[TestDatabaseConnectionResult]:
    if response.status_code == 200:
        response_200 = TestDatabaseConnectionResult.from_dict(response.json())

        return response_200
    if client.raise_on_unexpected_status:
        raise errors.UnexpectedStatus(response.status_code, response.content)
    else:
        return None


def _build_response(
    *, client: Union[AuthenticatedClient, Client], response: httpx.Response
) -> Response[TestDatabaseConnectionResult]:
    return Response(
        status_code=HTTPStatus(response.status_code),
        content=response.content,
        headers=response.headers,
        parsed=_parse_response(client=client, response=response),
    )


def sync_detailed(
    *,
    client: AuthenticatedClient,
    body: TestDatabaseConnectionPayload,
) -> Response[TestDatabaseConnectionResult]:
    """Test database connection

     Test a database connection

    Args:
        body (TestDatabaseConnectionPayload):

    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.

    Returns:
        Response[TestDatabaseConnectionResult]
    """

    kwargs = _get_kwargs(
        body=body,
    )

    response = client.get_httpx_client().request(
        **kwargs,
    )

    return _build_response(client=client, response=response)


def sync(
    *,
    client: AuthenticatedClient,
    body: TestDatabaseConnectionPayload,
) -> Optional[TestDatabaseConnectionResult]:
    """Test database connection

     Test a database connection

    Args:
        body (TestDatabaseConnectionPayload):

    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.

    Returns:
        TestDatabaseConnectionResult
    """

    return sync_detailed(
        client=client,
        body=body,
    ).parsed


async def asyncio_detailed(
    *,
    client: AuthenticatedClient,
    body: TestDatabaseConnectionPayload,
) -> Response[TestDatabaseConnectionResult]:
    """Test database connection

     Test a database connection

    Args:
        body (TestDatabaseConnectionPayload):

    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.

    Returns:
        Response[TestDatabaseConnectionResult]
    """

    kwargs = _get_kwargs(
        body=body,
    )

    response = await client.get_async_httpx_client().request(**kwargs)

    return _build_response(client=client, response=response)


async def asyncio(
    *,
    client: AuthenticatedClient,
    body: TestDatabaseConnectionPayload,
) -> Optional[TestDatabaseConnectionResult]:
    """Test database connection

     Test a database connection

    Args:
        body (TestDatabaseConnectionPayload):

    Raises:
        errors.UnexpectedStatus: If the server returns an undocumented status code and Client.raise_on_unexpected_status is True.
        httpx.TimeoutException: If the request takes longer than Client.timeout.

    Returns:
        TestDatabaseConnectionResult
    """

    return (
        await asyncio_detailed(
            client=client,
            body=body,
        )
    ).parsed
