import React from "react";
import {
    InputProps as IProps,
    Box, Text
} from "@chakra-ui/react";

import styles from "./FullLikeBox.module.scss";
import RubIcon from "components/Icons/Rub";

interface InputProps extends IProps {
    price?: number | string
}

const FullLikeBox: React.FC<InputProps> = (props) => {
    return (
    <Box className={styles.likes_box}> 
        <RubIcon />
        <Text> {props.price} </Text>
    </Box>
    )
}

export default FullLikeBox;