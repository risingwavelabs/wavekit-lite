/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { MetricsStoreLabelMatcherList } from './MetricsStoreLabelMatcherList';
import type { MetricsStoreSpec } from './MetricsStoreSpec';
export type MetricsStore = {
    ID: number;
    name: string;
    createdAt: string;
    spec?: MetricsStoreSpec;
    defaultLabels?: MetricsStoreLabelMatcherList;
};

