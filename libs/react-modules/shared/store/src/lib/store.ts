/* eslint-disable @nx/enforce-module-boundaries */
import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit';
import {
  CATEGORIES_FEATURE_KEY,
  categoriesTempReducer,
  apiSlice,
  categoriesApiSlice,
} from '@react-modules/features/categories/data-access';


export const store = configureStore({
  reducer: {
    [CATEGORIES_FEATURE_KEY]: categoriesTempReducer,
    [apiSlice.reducerPath]: apiSlice.reducer,
    [categoriesApiSlice.reducerPath]: categoriesApiSlice.reducer,
  },
  middleware: (getDefaultMiddleware: any) => getDefaultMiddleware(),
  devTools: process.env.NODE_ENV !== 'production',
  enhancers: [],
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
