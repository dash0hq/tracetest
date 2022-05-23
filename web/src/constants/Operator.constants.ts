import {TCompareOperatorName, TCompareOperatorSymbol} from '../types/Operator.types';

export enum CompareOperator {
  EQUALS = 'EQUALS',
  NOTEQUALS = 'NOTEQUALS',
  LESSTHAN = 'LESSTHAN',
  GREATERTHAN = 'GREATERTHAN',
  GREATOREQUALS = 'GREATOREQUALS',
  LESSOREQUAL = 'LESSOREQUAL',
  CONTAINS = 'CONTAINS',
}

export const CompareOperatorNameMap: Record<CompareOperator, TCompareOperatorName> = {
  [CompareOperator.EQUALS]: 'equals',
  [CompareOperator.NOTEQUALS]: 'not equals',
  [CompareOperator.LESSTHAN]: 'less than',
  [CompareOperator.GREATERTHAN]: 'greater than',
  [CompareOperator.GREATOREQUALS]: 'greater or equals',
  [CompareOperator.LESSOREQUAL]: 'less or equals',
  [CompareOperator.CONTAINS]: 'contains',
};

export const CompareOperatorSymbolMap: Record<CompareOperator, TCompareOperatorSymbol> = {
  [CompareOperator.EQUALS]: '=',
  [CompareOperator.LESSTHAN]: '<',
  [CompareOperator.GREATERTHAN]: '>',
  [CompareOperator.NOTEQUALS]: '!=',
  [CompareOperator.GREATOREQUALS]: '>=',
  [CompareOperator.LESSOREQUAL]: '<=',
  [CompareOperator.CONTAINS]: 'contains',
};

export enum PseudoSelector {
  FIRST = ':first',
  LAST = ':last',
  NTH = ':nth',
  ALL = ':all',
}
