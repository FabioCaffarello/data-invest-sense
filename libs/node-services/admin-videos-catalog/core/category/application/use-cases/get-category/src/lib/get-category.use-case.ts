import { IUseCase } from '@admin-videos-catalog/core/shared/application/use-case';
import { NotFoundError } from '@admin-videos-catalog/core/shared/errors';
import { Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid';
import { Category } from '@admin-videos-catalog/core/category/entity';
import { ICategoryRepository } from '@admin-videos-catalog/core/category/repository';
import { CategoryOutput, CategoryOutputMapper } from '@admin-videos-catalog/core/category/application/use-cases/common';

export class GetCategoryUseCase
  implements IUseCase<GetCategoryInput, GetCategoryOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: GetCategoryInput): Promise<GetCategoryOutput> {
    const uuid = new Uuid(input.id);
    const category = await this.categoryRepo.findById(uuid);
    if (!category) {
      throw new NotFoundError(input.id, Category);
    }

    return CategoryOutputMapper.toOutput(category);
  }
}

export type GetCategoryInput = {
  id: string;
};

export type GetCategoryOutput = CategoryOutput;
