// eslint-disable-next-line @typescript-eslint/no-unused-vars
// import './app.module.css';
import { Box, ThemeProvider } from '@mui/system';
import { Header, Layout } from '@react-modules/shared/ui-components';
import { darkTheme, lightTheme } from '../config/theme';
import { Routes, Route } from 'react-router-dom';
import { CategoryList, CategoryCreate, CategoryEdit } from '@react-modules/features/categories/ui';
import { Typography } from '@mui/material';
import { SnackbarProvider } from 'notistack';


export function App() {
  // return <Button variant='contained'>Hello world</Button>;
  return (
    <ThemeProvider theme={darkTheme}>
      <SnackbarProvider
      autoHideDuration={2000}
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        maxSnack={3}
      >

        <Box component='main'
          sx={{
            height: '100vh',
            backgroundColor: (theme) => theme.palette.background.default,
          }}
        >
          <Header />
          <Layout>
            <Routes>
              <Route path='/' element={<CategoryList />} />
              <Route path='/categories' element={<CategoryList />}/>
              <Route path='/categories/create' element={<CategoryCreate />} />
              <Route path='/categories/edit/:id' element={<CategoryEdit />} />
              <Route path='*' element={
                <Box sx={{ color: 'white' }}>
                  <Typography variant='h1'>404</Typography>
                  <Typography variant='h2'>Page not found</Typography>
                </Box>
              } />
            </Routes>
          </Layout>
        </Box>
      </SnackbarProvider>
    </ThemeProvider>
  );
}

export default App;

// if (import.meta.vitest) {P
//   // add tests related to your file here
//   // For more information please visit the Vitest docs site here: https://vitest.dev/guide/in-source.html

//   const { it, expect, beforeEach } = import.meta.vitest;
//   let render: any;

//   beforeEach(async () => {
//     render = (await import('@testing-library/react')).render;
//   });

//   it('should render successfully', () => {
//     const { baseElement } = render(<App />);
//     expect(baseElement).toBeTruthy();
//   });

//   it('should have a greeting as the title', () => {
//     const { getByText } = render(<App />);
//     expect(getByText(/Welcome products-front-end-midiaflix/gi)).toBeTruthy();
//   });
// }
