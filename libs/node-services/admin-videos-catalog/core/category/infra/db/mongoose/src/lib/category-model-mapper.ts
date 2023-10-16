import { EntityValidationError } from "@admin-videos-catalog/core/shared/validators";
import { Uuid } from "@admin-videos-catalog/core/shared/value-objects/uuid";
import { Category } from "@admin-videos-catalog/core/category/entity";
import { CategoryDocument } from "./category.model";


export class CategoryModelMapper {
  // static toModel(entity: Category): CategoryModel {
  //   return CategoryModel.build({
  //     category_id: entity.category_id.id,
  //     name: entity.name,
  //     description: entity.description,
  //     is_active: entity.is_active,
  //     created_at: entity.created_at,
  //   });
  // }

  // static toEntity(model: CategoryDocument): Category {
  //   const category = new Category({
  //     category_id: new Uuid(model.category_id),
  //     name: model.name,
  //     description: model.description,
  //     is_active: model.is_active,
  //     created_at: model.created_at,
  //   });
  //   // TODO: Validate entity
  //   // category.validate();
  //   if (category.notification.hasErrors()) {
  //     throw new EntityValidationError(category.notification.toJSON());
  //   }
  //   return category;
  // }
}
