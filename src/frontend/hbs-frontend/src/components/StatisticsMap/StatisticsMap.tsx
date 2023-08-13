import { Box, Text } from "@chakra-ui/react";
import React from "react";

import styles from "./Statistics.module.scss";
import { AllAvgQueryTimeResp, AllAvgServiceTimeResp, PopularHotelsResp } from "postAPI";
import StatisticsAvgServiceTimeCard from "components/StatisticsAvgServiceTimeCard/StatisticsAvgServiceTimeCard";
import StatisticsAvgQueryTimeCard from "components/StatisticsAvgQueryTimeCard/StatisticsAvgQueryTimeCard";
import PopularHotelCard from "components/PopulatHotelCard/PopularHotelCard";


interface StatisticsBoxProps {
    searchQuery?: string
    getCall: () => Promise<AllAvgServiceTimeResp>
    getCallQuery: () => Promise<AllAvgQueryTimeResp>
    getCallPopularHotels: () => Promise<PopularHotelsResp>
}

type State = {
    avgServiceTime?: any
    avgQueryTime?: any
    popularHotels?: any
}


class StatisticsMap extends React.Component<StatisticsBoxProps, State> {
    constructor(props) {
        super(props);
        this.state = {
            avgServiceTime: [],
            avgQueryTime: [],
            popularHotels: []
        }
    }

    async getAllAvgServiceTime() {
        var data = await this.props.getCall();
        if (data.status === 200)
            this.setState({avgServiceTime: data.content})
    }

    async getAllAvgQueryTime() {
        var data = await this.props.getCallQuery();
        if (data.status === 200)
            this.setState({avgQueryTime: data.content})
    }

    async getPopularHotels() {
        var data = await this.props.getCallPopularHotels()
        if (data.status === 200)
            this.setState({popularHotels: data.content})
    }

    componentDidMount() {
        this.getAllAvgServiceTime();
        this.getAllAvgQueryTime();
        this.getPopularHotels();
    }

    render() {
        return (
            <Box className={styles.map_box}>
                <Text className={styles.title_box}>
                    3 самых популярных отеля
                </Text>
                {this.state.popularHotels.map(item => <PopularHotelCard {...item} key={item.id}/>)}

                <Text className={styles.title_box}>
                    Общие характеристики по сервисам
                </Text>
                {this.state.avgServiceTime.map(item => <StatisticsAvgServiceTimeCard {...item} key={item.id}/>)}
                
                <Text className={styles.title_box}>
                    Общие характеристики по запросам
                </Text>
                {this.state.avgQueryTime.map(item => <StatisticsAvgQueryTimeCard {...item} key={item.id}/>)}
            </Box>
        )
    }
}

export default React.memo(StatisticsMap);