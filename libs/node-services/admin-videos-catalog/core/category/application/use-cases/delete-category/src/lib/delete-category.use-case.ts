import { IUseCase } from '@admin-videos-catalog/core/shared/application/use-case';
import { Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid';
import { ICategoryRepository } from '@admin-videos-catalog/core/category/repository';

export class DeleteCategoryUseCase
  implements IUseCase<DeleteCategoryInput, DeleteCategoryOutput>
{
  constructor(private categoryRepo: ICategoryRepository) {}

  async execute(input: DeleteCategoryInput): Promise<DeleteCategoryOutput> {
    const uuid = new Uuid(input.id);
    await this.categoryRepo.delete(uuid);
  }
}

export type DeleteCategoryInput = {
  id: string;
};

type DeleteCategoryOutput = void;
