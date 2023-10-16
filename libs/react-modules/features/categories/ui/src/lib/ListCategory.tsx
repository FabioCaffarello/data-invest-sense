import DeleteIcon from "@mui/icons-material/Delete";
import { Button, IconButton, Typography } from '@mui/material';
import { Box } from '@mui/system';
import { DataGrid, GridColDef, GridRenderCellParams, GridRowsProp, GridToolbar } from '@mui/x-data-grid';
import { deleteCategory, selectCategoriesEntitiesTemp, useGetCategoriesQuery } from '@react-modules/features/categories/data-access';
import { useAppDispatch, useAppSelector } from '@react-modules/shared/hooks';
import { useSnackbar } from "notistack";
import { Link } from 'react-router-dom';



export const CategoryList = () => {
  const { data, isFetching, error } = useGetCategoriesQuery();


  const categories = useAppSelector(selectCategoriesEntitiesTemp);
  const dispatch = useAppDispatch();
  const { enqueueSnackbar } = useSnackbar();

  const slotProps = {
    toolbar: {
      showQuickFilter: true,
    },
  }

  // use categories to create rows
  const rows: GridRowsProp = categories.map((category) => ({
    id: category.id,
    name: category.name,
    description: category.description,
    isActive: category.is_active,
    createdAt: new Date(category.created_at).toLocaleDateString('en-US')
  }));

  const columns: GridColDef[] = [
    {
      field: 'name',
      headerName: 'Name',
      flex: 1,
      renderCell: renderNameCell
    },
    {
      field: 'isActive',
      headerName: 'Active',
      flex: 1,
      type: 'boolean',
      renderCell: renderIsActiveCell,
    },
    {
      field: 'createdAt',
      headerName: 'Created At',
      flex: 1
    },
    {
      field: 'id',
      headerName: 'Actions',
      type: 'string',
      flex: 1,
      renderCell: renderActionsCell,
    },
  ];

  function renderNameCell(params: GridRenderCellParams) {
    return (
      <Link
        style={{ textDecoration: 'none' }}
        to={`/categories/edit/${params.id}`}
      >
        <Typography color="primary">{params.value}</Typography>
      </Link>
    )
  }

  function handleDeleteCategory(id: string) {
    dispatch(deleteCategory(id));
    enqueueSnackbar('Category deleted successfully', { variant: 'warning' });
  }

  function renderActionsCell(params: GridRenderCellParams) {
    return (
      <IconButton
        color="secondary"
        onClick={() => handleDeleteCategory(params.value)}
        aria-label="delete"
      >
        <DeleteIcon />
      </IconButton>
    )
  }

  function renderIsActiveCell(rowData: GridRenderCellParams) {
    return (
      <Typography color={rowData.value ? 'primary' : 'secondary'}>
        {rowData.value ? 'Active' : 'Inactive'}
      </Typography>
    )
  }

  return (
    <Box maxWidth="lg" sx={{ mt:4, mb:4 }}>
      <Box display="flex" justifyContent="flex-end">
        <Button
          variant="contained"
          color="secondary"
          component={Link}
          to="/categories/create"
          style={{ marginBottom: '1rem' }}
        >
          New Category
        </Button>
      </Box>

      <Box sx={{ display: "flex", height: 600 }}>
        <DataGrid
          rows={rows}
          columns={columns}
          initialState={{
            pagination: {
              paginationModel: { page: 0, pageSize: 2 },
            },
            filter: {
              filterModel: {
                items: [],
                quickFilterValues: [],
              },
            },
          }}
          pageSizeOptions={[2, 5, 10]}
          checkboxSelection={false}
          slots={{ toolbar: GridToolbar }}
          slotProps={slotProps}
          disableColumnFilter={true}
          disableColumnSelector={true}
          disableDensitySelector={true}
          disableRowSelectionOnClick={true}
        />
      </Box>
    </Box>
  );
};
