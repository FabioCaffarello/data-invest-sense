import { Transform } from 'class-transformer';
import { CategoryOutput } from '@admin-videos-catalog/core/category/application/use-cases/common';
import { ListCategoriesOutput } from '@admin-videos-catalog/core/category/application/use-cases/list-categories';
import { CollectionPresenter } from '@admin-videos-catalog/nest-modules/shared-module';

export class CategoryPresenter {
  id: string;
  name: string;
  description: string | null;
  is_active: boolean;
  @Transform(({ value }: { value: Date }) => value.toISOString())
  created_at: Date;

  constructor(output: CategoryOutput) {
    this.id = output.id;
    this.name = output.name;
    this.description = output.description;
    this.is_active = output.is_active;
    this.created_at = output.created_at;
  }
}

export class CategoryCollectionPresenter extends CollectionPresenter {
  data: CategoryPresenter[];

  constructor(output: ListCategoriesOutput) {
    const { items, ...paginationProps } = output;
    super(paginationProps);
    this.data = items.map((i) => new CategoryPresenter(i));
  }
}
