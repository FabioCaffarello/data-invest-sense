import {
  PaginationOutput,
  PaginationOutputMapper,
  IUseCase
} from '@admin-videos-catalog/core/shared/application/use-case';
import { SortDirection } from '@admin-videos-catalog/core/shared/repository';
import {
  CategoryFilter,
  CategorySearchParams,
  CategorySearchResult,
  ICategoryRepository,
} from '@admin-videos-catalog/core/category/repository';
import { CategoryOutput, CategoryOutputMapper } from '@admin-videos-catalog/core/category/application/use-cases/common';

export class ListCategoriesUseCase
  implements IUseCase<ListCategoriesInput, ListCategoriesOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: ListCategoriesInput): Promise<ListCategoriesOutput> {
    const params = new CategorySearchParams(input);
    const searchResult = await this.categoryRepo.search(params);
    return this.toOutput(searchResult);
  }

  private toOutput(searchResult: CategorySearchResult): ListCategoriesOutput {
    const { items: _items } = searchResult;
    const items = _items.map((i) => {
      return CategoryOutputMapper.toOutput(i);
    });
    return PaginationOutputMapper.toOutput(items, searchResult);
  }
}

export type ListCategoriesInput = {
  page?: number;
  per_page?: number;
  sort?: string | null;
  sort_dir?: SortDirection | null;
  filter?: CategoryFilter | null;
};

export type ListCategoriesOutput = PaginationOutput<CategoryOutput>;
