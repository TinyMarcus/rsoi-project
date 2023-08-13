import { Box, HStack, Image, Text, VStack } from "@chakra-ui/react";
import React from "react";

import { Hotel as HotelI } from "types/Hotel";

import StarBox from "components/Boxes/Star";
import FullLikeBox from "components/Boxes/FullLike";

import styles from "./HotelCard.module.scss";
import GetImageUrl from "postAPI/likes/Get";

import IngredientModel from "components/DateInput";
import {DateReservation as DateReservationT} from "types/DateReservation";
import {ReservationRequest as ReservationRequestT} from "types/ReservationRequest";
import PostReservation from "postAPI/likes/Post";


interface HotelProps extends HotelI {}

const HotelCard: React.FC<HotelProps> = (props) => {
  const [imageUrl, setImageUrl] = React.useState("https://media.discordapp.net/attachments/791290400086032437/1112498073478889502/default-fallback-image.png?width=720&height=480");

  async function getImageUrl() {
    var data = await GetImageUrl(props.hotel_uid);
    if (data.status === 200) {
      setImageUrl(data.content);
    }
  }

  async function putDateReservation(dataInfo: DateReservationT) {
    const reservationRequest: ReservationRequestT = { hotel_uid: props.hotel_uid,
      start_date: dataInfo.start_date, //.toLocaleDateString("en-CA"),
      end_date: dataInfo.end_date }; //.toLocaleDateString("en-CA") };

    await PostReservation(reservationRequest);
  }

  getImageUrl();

  return (
    <Box className={styles.main_box}>
      <Image src={imageUrl} className={styles.image_div} />

      <Box className={styles.info_box}>
        <VStack>
          <Box className={styles.title_box}>
            <Text>{props.name}</Text>
          </Box>

          <Box className={styles.description_box}>
            <Text>{props.country}</Text>
          </Box>
          <Box className={styles.description_box}>
            <Text>{props.city}</Text>
          </Box>
          <Box className={styles.description_box}>
            <Text>{props.address}</Text>
          </Box>
        </VStack>

        <HStack>
          <StarBox duration={props.stars} />
          <FullLikeBox price={props.price} />
        </HStack>

        <IngredientModel putCallback={(data: DateReservationT) => putDateReservation(data)}/>

      </Box>
    </Box>
  );
};

export default HotelCard;
