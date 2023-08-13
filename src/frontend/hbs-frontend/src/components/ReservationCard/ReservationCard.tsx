import { Box, HStack, Image, Text, VStack } from "@chakra-ui/react";
import React from "react";

import { Reservation as ReservationI } from "types/Reservation";

import StarBox from "components/Boxes/Star";
import FullLikeBox from "components/Boxes/FullLike";

import styles from "./ReservationCard.module.scss";
import GetImageUrl from "postAPI/likes/Get";
import CancelReservation from "postAPI/likes/Cancel";
import { statusItems } from "./items";
import RoundButton from "components/RoundButton/RoundButton";


interface ReservationProps extends ReservationI {}

const ReservationCard: React.FC<ReservationProps> = (props) => {
  const [imageUrl, setImageUrl] = React.useState("https://media.discordapp.net/attachments/791290400086032437/1112498073478889502/default-fallback-image.png?width=720&height=480");
  const [status, setStatus] = React.useState(statusItems[props.status]);

  async function getImageUrl() {
    var data = await GetImageUrl(props.hotel.hotel_uid);
    if (data.status === 200) {
      setImageUrl(data.content);
    }
  }

  async function submit(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
    await CancelReservation(props.reservation_uid);
    setStatus(statusItems['CANCELED']);
  }

  getImageUrl();

  return (
    <Box className={styles.main_box}>
    <Image src={imageUrl} className={styles.image_div} />

    <Box className={styles.info_box}>
        <VStack>
        <Box className={styles.title_box}>
            <Text>{props.hotel.name}</Text>
        </Box>

        <Box className={styles.description_box}>
            <Text>{props.hotel.fullAddress}</Text>
        </Box>
        <Box className={styles.description_box}>
            <Text>Заезд: {new Date(props.start_date).toLocaleDateString("en-CA")}</Text>
        </Box>
        <Box className={styles.description_box}>
            <Text>Выезд: {new Date(props.end_date).toLocaleDateString("en-CA")}</Text>
        </Box>
        { status === 'Оплачено' &&
        <Box className={styles.description_box}>
            <Text>Статус: { status }</Text>
        </Box> }
        { status === 'Отменено' &&
        <Box className={styles.description_box}>
            <Text>Статус: { status }</Text>
        </Box> }

        <HStack>
            <StarBox duration={props.hotel.stars} />
            <FullLikeBox price={props.payment.price} />
        </HStack>
        </VStack>

        { (status === 'Оплачено') &&
            <RoundButton type="submit" onClick={event => submit(event)}>
                Отменить
            </RoundButton> 
        }
    </Box>
    </Box>
  );
};

export default ReservationCard;
