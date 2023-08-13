import { Box } from "@chakra-ui/react";
import GetMyReservations from "postAPI/likes/GetMyReservations";
import React from "react";

import styles from "./ReservationsPage.module.scss";
import ReservationMap from "components/ReservationMap/ReservationMap";

interface ReservationsProps {}

const ReservationsPage: React.FC<ReservationsProps> = () => {
  return (
    <Box className={styles.main_box}>
      <ReservationMap getCall={GetMyReservations} />
    </Box>
  );
};

export default ReservationsPage;
