import * as React from "react";
import theme from "./styles/extendTheme";
import { ChakraProvider, Box, Container } from "@chakra-ui/react";
import { Routes, Route } from "react-router";
import { BrowserRouter } from "react-router-dom";

import Login from "pages/Login";
import SignUp from "pages/Signup";
import RecipeInfoPage from "pages/Recipe/RecipeInfo";
import AllHotelsPage from "pages/Recipe/AllHotels";
import UsersPage from "pages/Users/UsersPage";
// import StatisticsPage from "pages/Recipe/AllStatistics";
import StatisticsPage from "pages/Statistics/StatisticsPage"
import LikedRecipesPage from "pages/Recipe/AuthorReservations";
import CategoryPage from "pages/Category";

import SearchContextProvider from "context/Search";
import { HeaderRouter } from "components/Header";


interface HomeProps {}
const Home: React.FC<HomeProps> = () => {
  return (
    <Box backgroundColor="bg" h="auto">
      <Container maxW="1000px" minH="95%"
        display="flex" 
        paddingX="0px" paddingY="30px"  
        alignSelf="center" justifyContent="center"
        textStyle="body"
      >
        <Routing />
      </Container>
    </Box>
  );
};

function Routing() {
  return <BrowserRouter>
    <Routes>
      <Route path="/" element={<AllHotelsPage/>}/>
      <Route path="/me/reservations" element={<LikedRecipesPage/>}/>
      <Route path="/statistics" element={<StatisticsPage/>}/>

      <Route path="/accounts/:login/recipes" element={<StatisticsPage/>}/>
      <Route path="/accounts/:login/likes" element={<LikedRecipesPage/>}/>

      <Route path="/authorize" element={<Login/>}/>
      <Route path="/register" element={<SignUp/>}/>

      <Route path="/users" element={<UsersPage />}/>

      <Route path="*" element={<NotFound />}/>
    </Routes>
  </BrowserRouter>
}

function NotFound () {
  return <h1>Page not Found</h1>
}

export const App = () => {
  return (    
    <ChakraProvider theme={theme}>
    <SearchContextProvider>
      <HeaderRouter/>
      <Home />
    </SearchContextProvider>
    </ChakraProvider>
  )
};
