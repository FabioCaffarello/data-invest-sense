/* eslint-disable @typescript-eslint/no-empty-interface */
import {
  ISearchableRepository
} from '@admin-videos-catalog/core/shared/repository'
import { SearchParams, SearchResult } from '@admin-videos-catalog/core/shared/repository'
import { Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid'
import { Category } from '@admin-videos-catalog/core/category/entity'

export type CategoryFilter = string;

export class CategorySearchParams extends SearchParams<CategoryFilter> {}

export class CategorySearchResult extends SearchResult<Category> {}

export interface ICategoryRepository
  extends ISearchableRepository<
    Category,
    Uuid,
    CategoryFilter,
    CategorySearchParams,
    CategorySearchResult
  > {}
