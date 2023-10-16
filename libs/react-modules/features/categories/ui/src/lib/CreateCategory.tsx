import { Paper, Typography } from '@mui/material';
import { Box } from '@mui/system';
import { CategoryFrom } from '@react-modules/features/categories/components';
import { CategoriesEntity, createCategory } from '@react-modules/features/categories/data-access';
import { useAppDispatch } from '@react-modules/shared/hooks';
import { useSnackbar } from 'notistack';
import { useState } from 'react';


export function CategoryCreate() {
  const [isdisabled, setIsdisabled] = useState(false);
  const [categoryState, setCategoryState] = useState<CategoriesEntity>({
    id: "",
    name: "",
    description: "",
    is_active: false,
    created_at: "",
    updated_at: "",
    deleted_at: "",
  });
  const dispatch = useAppDispatch();
  const { enqueueSnackbar } = useSnackbar();

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    dispatch(createCategory(categoryState))
    enqueueSnackbar('Category created successfully', { variant: 'success' });
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
            <Typography variant='h4'>Create Category</Typography>
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
