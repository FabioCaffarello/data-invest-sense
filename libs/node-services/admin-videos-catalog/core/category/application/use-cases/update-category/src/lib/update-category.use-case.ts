import { IUseCase } from '@admin-videos-catalog/core/shared/application/use-case';
import { NotFoundError } from '@admin-videos-catalog/core/shared/errors';
import { EntityValidationError } from '@admin-videos-catalog/core/shared/validators';
import { Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid';
import { Category } from '@admin-videos-catalog/core/category/entity';
import { ICategoryRepository } from '@admin-videos-catalog/core/category/repository';
import {
  CategoryOutput,
  CategoryOutputMapper,
} from '@admin-videos-catalog/core/category/application/use-cases/common';
import { UpdateCategoryInput } from './update-category.input';

export class UpdateCategoryUseCase
  implements IUseCase<UpdateCategoryInput, UpdateCategoryOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: UpdateCategoryInput): Promise<UpdateCategoryOutput> {
    const uuid = new Uuid(input.id);
    const category = await this.categoryRepo.findById(uuid);

    if (!category) {
      throw new NotFoundError(input.id, Category);
    }

    input.name && category.changeName(input.name);

    if (input.description !== undefined) {
      category.changeDescription(input.description);
    }

    if (input.is_active === true) {
      category.activate();
    }

    if (input.is_active === false) {
      category.deactivate();
    }

    if (category.notification.hasErrors()) {
      throw new EntityValidationError(category.notification.toJSON());
    }

    await this.categoryRepo.update(category);

    return CategoryOutputMapper.toOutput(category);
  }
}

export type UpdateCategoryOutput = CategoryOutput;
