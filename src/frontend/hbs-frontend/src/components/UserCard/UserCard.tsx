import { Box, HStack, Image, Link, Text, VStack } from "@chakra-ui/react";
import React from "react";

import { Account as AccountI } from "types/Account";

import FullLikeBox from "components/Boxes/FullLike/FullLikeBox";
import RoleBox from "components/Boxes/Role/RoleBox";
import LoginBox from "components/LoginBox";
import RecipeBox from "components/RecipeBox";
import user from "img/user.png";

import styles from "./UserCard.module.scss";

interface UserProps extends AccountI {
    role: string
}

const UserCard: React.FC<UserProps> = (props) => {
  const pathRecipes = "/accounts/" + props.username + "/recipes";
  const pathLiked = "/accounts/" + props.username + "/likes";

  return (
    <Box className={styles.main_box}>
      <HStack>
        <VStack>
          <Image src={user} />

          <VStack className={styles.role}>
            <Text>Роль</Text>
            <RoleBox login={props.username} role={props.role}/>
          </VStack>
        </VStack>

        <VStack className={styles.info}>
          <Text>Логин</Text>
          <LoginBox login={props.username} className={styles.login} />

          <Link href={pathRecipes}>
            <RecipeBox data={"Отели"} className={styles.recipes} />
          </Link>

          <Link href={pathLiked}>
            {/* <FullLikeBox likesNum={"Понравилось"} /> */}
          </Link>
        </VStack>
      </HStack>
    </Box>
  );
};

export default UserCard;
