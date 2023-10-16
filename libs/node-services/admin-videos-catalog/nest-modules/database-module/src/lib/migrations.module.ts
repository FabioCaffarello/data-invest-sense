import { Module } from '@nestjs/common';
import { ConfigModule } from '@admin-videos-catalog/nest-modules/config-module';
import { DatabaseModule } from './database.module';

@Module({
  imports: [ConfigModule.forRoot(), DatabaseModule],
})
export class MigrationsModule {}
