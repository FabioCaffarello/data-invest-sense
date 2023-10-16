import { Paper, Typography } from '@mui/material';
import { Box } from '@mui/system';
import { CategoryFrom } from '@react-modules/features/categories/components';
import { CategoriesEntity, selectCategoryEntityById, updateCategory } from '@react-modules/features/categories/data-access';
import { useAppDispatch, useAppSelector } from '@react-modules/shared/hooks';
import { useSnackbar } from 'notistack';
import { useState } from 'react';
import { useParams } from 'react-router-dom';


export function CategoryEdit() {
  const id = useParams().id || "";
  const [isdisabled, setIsdisabled] = useState(false);
  const category = useAppSelector((state) => selectCategoryEntityById(state, id));
  const [categoryState, setCategoryState] = useState<CategoriesEntity>(category);
  const dispatch = useAppDispatch();
  const { enqueueSnackbar } = useSnackbar();

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    dispatch(updateCategory(categoryState))
    enqueueSnackbar('Category updated successfully', { variant: 'success' });
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setCategoryState({ ...categoryState, [name]: value })
  };

  const handleToggle = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = event.target;
    setCategoryState({ ...categoryState, [name]: checked })
  };

  return (
    <Box>
      <Paper>
        <Box p={2}>
          <Box mb={2}>
            <Typography variant='h4'>Edit Category</Typography>
          </Box>
        </Box>

      <CategoryFrom
        category={categoryState}
        isdisabled={isdisabled}
        isLoading={false}
        handleSubmit={handleSubmit}
        handleChange={handleChange}
        handleToggle={handleToggle}
      />
      </Paper>
    </Box>
  );
}
