import React from "react";
import {
    InputProps as IProps,
    Box, Text
} from "@chakra-ui/react";
import StarIcon from "components/Icons/Star";

import styles from "./StarBox.module.scss";

interface InputProps extends IProps {
    duration?: number
}

const StarBox: React.FC<InputProps> = (props) => {
    var stringDuration = ""

    if (!props.duration)
        stringDuration = "---"
    else if (props.duration == 1)
        stringDuration += props.duration + " звезда"
    else if (props.duration > 1 && props.duration < 5)
        stringDuration += props.duration + " звезды"
    else
        stringDuration += props.duration + " звёзд"

    return (
    <Box className={styles.round_box_tiny}> 
        <Box> <StarIcon /> </Box>
        <Text> {stringDuration} </Text>
    </Box>
    )
}

export default StarBox;