import { Module } from '@nestjs/common';
import { ConfigModule } from '@admin-videos-catalog/nest-modules/config-module';
import { DatabaseModule } from '@admin-videos-catalog/nest-modules/database-module';
import { CategoriesModule } from '@admin-videos-catalog/nest-modules/features/categories-module';
import { SharedModule } from '@admin-videos-catalog/nest-modules/shared-module';

@Module({
  imports: [ConfigModule.forRoot(), DatabaseModule, CategoriesModule, SharedModule],
})
export class AppModule {}
