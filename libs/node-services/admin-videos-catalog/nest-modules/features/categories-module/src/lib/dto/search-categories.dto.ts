import { ListCategoriesInput } from '@admin-videos-catalog/core/category/application/use-cases/list-categories';
import { SortDirection } from '@admin-videos-catalog/core/shared/repository';

export class SearchCategoriesDto implements ListCategoriesInput {
  page?: number;
  per_page?: number;
  sort?: string;
  sort_dir?: SortDirection;
  filter?: string;
}
