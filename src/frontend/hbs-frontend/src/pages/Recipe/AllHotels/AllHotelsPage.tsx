import { Box } from "@chakra-ui/react";
import CategoryMap from "components/CategoryMap";
import { SearchContext } from "context/Search";
import GetCategories from "postAPI/categories/GetAll";
import GetHotels from "postAPI/recipes/GetAll";
import React, { useContext } from "react";
import RecipeMap from "../../../components/RecipeMap/RecipeMap";

import styles from "./AllHotelsPage.module.scss";

interface AllRecipesProps {}

const AllHotelsPage: React.FC<AllRecipesProps> = (props) => {
  const searchContext = useContext(SearchContext);

  if (localStorage.getItem("token") == null) {
    window.location.href = "/authorize";
    return (<Box></Box>);
  }

  return (
    <Box className={styles.main_box}>
      {/* <CategoryMap getCall={GetCategories}/> */}
      <RecipeMap searchQuery={searchContext.query} getCall={GetHotels}/>
    </Box>
  );
};

export default AllHotelsPage;
