/* eslint-disable */
export default {
  displayName:
    'node-services-admin-videos-catalog-core-category-infra-db-sequelize',
  preset: '../../../../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../../../../coverage/libs/node-services/admin-videos-catalog/core/category/infra/db/sequelize',
    setupFilesAfterEnv: ["../../../../../../expect-helpers.ts"],
};
