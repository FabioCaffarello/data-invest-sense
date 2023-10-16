import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

@Schema({
  timestamps: true
})
export class CategoryModelProps {
  @Prop() category_id: string;
  @Prop() name: string;
  @Prop() description: string;
  @Prop() is_active: boolean;
  @Prop() created_at: Date;
}

export type CategoryDocument = CategoryModelProps & Document;

export const CategorySchema = SchemaFactory.createForClass(CategoryModelProps);
