import { Box, Container } from "@mui/material";

/* eslint-disable-next-line */
export interface LayoutProps {
  children: React.ReactNode;
}

export function Layout(props: LayoutProps) {
  return (
    <Box>
      <Container
        maxWidth='lg'
        sx={{
          color: 'white',
          mt:4,
          mb:4
        }}
      >
          {props.children}
      </Container>
    </Box>
  );
}

export default Layout;
