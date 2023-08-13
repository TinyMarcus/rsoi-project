import { Box, Text, VStack } from "@chakra-ui/react";
import React from "react";
import { AvgServiceTime as AvgServiceTimeI } from "types/AvgServiceTime";

import styles from "./StatisticsAvgServiceTimeCard.module.scss";


interface StatisticsAvgServiceTimeProps extends AvgServiceTimeI {}


const StatisticsAvgServiceTimeCard: React.FC<StatisticsAvgServiceTimeProps> = (props) => {
    return (
        <Box className={styles.main_box}>
            <Box className={styles.info_box}>
                <VStack>
                    <Box className={styles.title_box}>
                        <Text>Сервис: {props.service}</Text>
                    </Box>

                    <Box className={styles.description_box}>
                        <Text>Количество запросов: {props.num} шт</Text>
                    </Box>

                    <Box className={styles.description_box}>
                        <Text>Время работы: {props.avg_time.toFixed(3)} мс</Text>
                    </Box>
                </VStack>
            </Box>
        </Box>
    )
};

export default StatisticsAvgServiceTimeCard;