/* eslint-disable @nx/enforce-module-boundaries */
import {
  createSlice,
} from '@reduxjs/toolkit';
import { RootState } from '@react-modules/shared/store';
import { apiSlice } from './categories-api.slice';
import { Results } from '@react-modules/features/categories/types';

export const CATEGORIES_FEATURE_KEY = 'categories';

const category: CategoriesEntity = {
  id: '0ne68ddd-0ne6-0ne6-0ne6-0ne68ddd0ne6',
  name: 'Category 1',
  description: 'Category 1 description',
  is_active: true,
  deleted_at: null,
  created_at: '2023-10-16T10:08:09+0000',
  updated_at: '2023-10-16T10:08:09+0000',
};

const endpointUrl = '/categories';
export const categoriesApiSlice = apiSlice.injectEndpoints({
  endpoints: ({ query }) => ({
    getCategories: query<Results, void>({
      query: () => `${endpointUrl}`,
      providesTags: ['Categories'],
    })
  })
});


export const initialState = [
  category,
  { ...category, id: '0ne68ddd-0ne6-0ne6-0ne6-0ne68ddd0ne7', name: 'Category 2' },
  { ...category, id: '0ne68ddd-0ne6-0ne6-0ne6-0ne68ddd0ne8', name: 'Category 3' },
  { ...category, id: '0ne68ddd-0ne6-0ne6-0ne6-0ne68ddd0ne9', name: 'Category 4' },
]

export interface CategoriesEntity {
  id: string;
  name: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  deleted_at: null | string;
  description: null | string;
}


export const categoriesTempSlice = createSlice({
  name: CATEGORIES_FEATURE_KEY,
  initialState,
  reducers: {
    createCategory(state, action) {
      state.push(action.payload);
    },
    updateCategory(state, action) {
      const index = state.findIndex(
        (category: CategoriesEntity) => category.id === action.payload.id
      );
      state[index] = action.payload;
    },
    deleteCategory(state, action) {
      const index = state.findIndex(
        (category: CategoriesEntity) => category.id === action.payload.id
      );
      state.splice(index, 1);
    },
  },
});

export const categoriesTempReducer = categoriesTempSlice.reducer;
export const { createCategory, updateCategory, deleteCategory } = categoriesTempSlice.actions;

// Selectors
export const selectCategoriesEntitiesTemp = (state: RootState) => state[CATEGORIES_FEATURE_KEY];
// Select Category by id
export const selectCategoryEntityById = (state: RootState, id: string) => {
  const category = state[CATEGORIES_FEATURE_KEY].find((category: CategoriesEntity) => category.id === id);
  return category || {
    id: '',
    name: '',
    is_active: false,
    created_at: '',
    updated_at: '',
    deleted_at: null,
    description: null,
  }
}


export const {
  useGetCategoriesQuery,
} = categoriesApiSlice
