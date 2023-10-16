import { OmitType } from '@nestjs/mapped-types';
import { UpdateCategoryInput } from '@admin-videos-catalog/core/category/application/use-cases/update-category';

export class UpdateCategoryInputWithoutId extends OmitType(
  UpdateCategoryInput,
  ['id'] as const,
) {}

export class UpdateCategoryDto extends UpdateCategoryInputWithoutId {}
