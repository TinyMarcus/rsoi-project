import { Box, Text, VStack, HStack } from "@chakra-ui/react";
import React from "react";
import { PopularHotel as PopularHotelI } from "types/PopularHotel";

import styles from "./PopularHotelCard.module.scss";

interface PopularHotelProps extends PopularHotelI {}

const PopularHotelCard: React.FC<PopularHotelProps> = (props) => {
    return (
        <Box className={styles.main_box}>
            <Box className={styles.info_box}>
                <VStack>
                    <Box className={styles.title_box}>
                        <Text>Отель: {props.name}</Text>
                    </Box>

                    <Box className={styles.description_box}>
                        <Text>Количество бронирований: {props.cnt} шт</Text>
                    </Box>
                </VStack>
            </Box>
        </Box>
    );
}

export default PopularHotelCard;