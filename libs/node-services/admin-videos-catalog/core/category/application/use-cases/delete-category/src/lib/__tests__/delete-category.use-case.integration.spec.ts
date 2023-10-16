import { NotFoundError } from '@admin-videos-catalog/core/shared/errors';
import { Uuid } from '@admin-videos-catalog/core/shared/value-objects/uuid';
import { setupSequelize } from '@admin-videos-catalog/core/shared/infra/testing';
import { Category } from '@admin-videos-catalog/core/category/entity';
import { CategorySequelizeRepository, CategoryModel } from '@admin-videos-catalog/core/category/infra/db/sequelize';
import { DeleteCategoryUseCase } from '../delete-category.use-case';

describe('DeleteCategoryUseCase Integration Tests', () => {
  let useCase: DeleteCategoryUseCase;
  let repository: CategorySequelizeRepository;

  setupSequelize({ models: [CategoryModel] });

  beforeEach(() => {
    repository = new CategorySequelizeRepository(CategoryModel);
    useCase = new DeleteCategoryUseCase(repository);
  });

  it('should throws error when entity not found', async () => {
    const uuid = new Uuid();
    await expect(() => useCase.execute({ id: uuid.id })).rejects.toThrow(
      new NotFoundError(uuid.id, Category)
    );
  });

  it('should delete a category', async () => {
    const category = Category.fake().aCategory().build();
    await repository.insert(category);
    await useCase.execute({
      id: category.category_id.id,
    });
    await expect(repository.findById(category.category_id)).resolves.toBeNull();
  });
});
