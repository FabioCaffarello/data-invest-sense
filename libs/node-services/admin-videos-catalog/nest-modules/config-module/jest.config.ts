/* eslint-disable */
export default {
  displayName: 'node-services-admin-videos-catalog-nest-modules-config-module',
  preset: '../../../../../jest.preset.js',
  testEnvironment: 'node',
  transform: {
    '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
  },
  moduleFileExtensions: ['ts', 'js', 'html'],
  coverageDirectory:
    '../../../../../coverage/libs/node-services/admin-videos-catalog/nest-modules/config-module',
};
