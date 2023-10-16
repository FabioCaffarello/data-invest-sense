import { NestFactory } from '@nestjs/core';
import { MigrationsModule } from '@admin-videos-catalog/nest-modules/database-module';
import { getConnectionToken } from '@nestjs/sequelize';
import { migrator } from '@admin-videos-catalog/core/shared/infra/db/sequelize';

async function bootstrap() {
  const app = await NestFactory.createApplicationContext(MigrationsModule, {
    logger: ['error'],
  });

  const sequelize = app.get(getConnectionToken());

  migrator(sequelize).runAsCLI();
}
bootstrap();
