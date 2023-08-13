import { Box } from "@chakra-ui/react";
import React from "react";
import HotelCard from "../HotelCard";
import { AllReservationsResp } from "postAPI";

import styles from "./ReservationMap.module.scss";
import ReservationCard from "components/ReservationCard/ReservationCard";

interface ReservationBoxProps {
    searchQuery?: string
    getCall: (title?: string) => Promise<AllReservationsResp>
}

type State = {
    postContent?: any
}

class ReservationMap extends React.Component<ReservationBoxProps, State> {
    constructor(props) {
        super(props);
        this.state = {
            postContent: []
        }
    }

    async getAll() {
        var data = await this.props.getCall(this.props.searchQuery)
        if (data.status === 200)
            this.setState({postContent: data.content});
    }

    componentDidMount() {
        this.getAll()
    }

    componentDidUpdate(prevProps) {
        if (this.props.searchQuery !== prevProps.searchQuery) {
            this.getAll()
        }
    }

    render() {
        return (
            <Box className={styles.map_box}>
                {this.state.postContent.map(item => <ReservationCard {...item} key={item.id}/>)}
            </Box>
        )
    }
}

export default React.memo(ReservationMap);
