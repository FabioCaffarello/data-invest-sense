import { NotFoundError } from '@admin-videos-catalog/core/shared/errors';
import { InvalidUuidError, Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid';
import { Category } from '@admin-videos-catalog/core/category/entity';
import { CategoryInMemoryRepository } from '@admin-videos-catalog/core/category/infra/db/in-memory';
import { DeleteCategoryUseCase } from '../delete-category.use-case';

describe('DeleteCategoryUseCase Unit Tests', () => {
  let useCase: DeleteCategoryUseCase;
  let repository: CategoryInMemoryRepository;

  beforeEach(() => {
    repository = new CategoryInMemoryRepository();
    useCase = new DeleteCategoryUseCase(repository);
  });

  it('should throws error when entity not found', async () => {
    await expect(() =>
      useCase.execute({ id: 'fake id'})
    ).rejects.toThrow(new InvalidUuidError());

    const uuid = new Uuid();

    await expect(() =>
      useCase.execute({ id: uuid.id})
    ).rejects.toThrow(new NotFoundError(uuid.id, Category));
  });

  it('should delete a category', async () => {
    const items = [new Category({ name: 'test 1' })];
    repository.items = items;
    await useCase.execute({
      id: items[0].category_id.id,
    });
    expect(repository.items).toHaveLength(0);
  });
});
