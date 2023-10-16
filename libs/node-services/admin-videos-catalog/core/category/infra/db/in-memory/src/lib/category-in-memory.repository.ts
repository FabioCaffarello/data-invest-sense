import { SortDirection } from '@admin-videos-catalog/core/shared/repository'
import { Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid'
import { InMemorySearchableRepository } from '@admin-videos-catalog/core/shared/infra/db/in-memory'
import { Category } from '@admin-videos-catalog/core/category/entity'
import { CategoryFilter, ICategoryRepository } from '@admin-videos-catalog/core/category/repository'


export class CategoryInMemoryRepository
  extends InMemorySearchableRepository<Category, Uuid>
  implements ICategoryRepository
{
  sortableFields: string[] = ["name", "created_at"];

  protected async applyFilter(
    items: Category[],
    filter: CategoryFilter
  ): Promise<Category[]> {
    if (!filter) {
      return items;
    }

    return items.filter((i) => {
      return i.name.toLowerCase().includes(filter.toLowerCase());
    });
  }
  getEntity(): new (...args: any[]) => Category {
    return Category;
  }

  protected applySort(
    items: Category[],
    sort: string | null,
    sort_dir: SortDirection | null
  ) {
    return sort
      ? super.applySort(items, sort, sort_dir)
      : super.applySort(items, "created_at", "desc");
  }
}
