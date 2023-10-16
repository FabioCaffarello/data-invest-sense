import { IUseCase } from '@admin-videos-catalog/core/shared/application/use-case';
import { EntityValidationError } from '@admin-videos-catalog/core/shared/validators';
import { Category } from '@admin-videos-catalog/core/category/entity';
import { ICategoryRepository } from '@admin-videos-catalog/core/category/repository';
import {
  CategoryOutput,
  CategoryOutputMapper,
} from '@admin-videos-catalog/core/category/application/use-cases/common';
import { CreateCategoryInput } from './create-category.input';

export class CreateCategoryUseCase
  implements IUseCase<CreateCategoryInput, CreateCategoryOutput>
{
  constructor(private readonly categoryRepo: ICategoryRepository) {}

  async execute(input: CreateCategoryInput): Promise<CreateCategoryOutput> {
    const entity = Category.create(input);

    if (entity.notification.hasErrors()) {
      throw new EntityValidationError(entity.notification.toJSON());
    }

    await this.categoryRepo.insert(entity);

    return CategoryOutputMapper.toOutput(entity);
  }
}

export type CreateCategoryOutput = CategoryOutput;
