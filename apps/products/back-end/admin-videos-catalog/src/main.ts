import { NestFactory } from '@nestjs/core';
import { AppModule } from './app/app.module';
import { applyGlobalConfig } from './app/global-config';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  applyGlobalConfig(app);

  await app.listen(3000);
}
bootstrap();
