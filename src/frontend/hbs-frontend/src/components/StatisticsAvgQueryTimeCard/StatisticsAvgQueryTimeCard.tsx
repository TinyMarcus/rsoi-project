import { Box, Text, VStack } from "@chakra-ui/react";
import React from "react";
import { AvgQueryServiceTime as AvgQueryServiceTimeI } from "types/AvgQueryServiceTime";

import styles from "./StatisticsAvgQueryTimeCard.module.scss";

interface StatisticsAvgQueryTimeProps extends AvgQueryServiceTimeI {}


const StatisticsAvgQueryTimeCard: React.FC<StatisticsAvgQueryTimeProps> = (props) => {
    return (
        <Box className={styles.main_box}>
            <Box className={styles.info_box}>
                <VStack>
                    <Box className={styles.title_box}>
                        <Text>Сервис: {props.service}</Text>
                    </Box>

                    <Box className={styles.description_box}>
                        <Text>Запрос: {props.action}</Text>
                    </Box>

                    <Box className={styles.description_box}>
                        <Text>Количество запросов: {props.num} шт</Text>
                    </Box>

                    <Box className={styles.description_box}>
                        <Text>Среднее время выполнения: {props.avg_time.toFixed(3)} мс</Text>
                    </Box>
                </VStack>
            </Box>
        </Box>
    )
};

export default StatisticsAvgQueryTimeCard;