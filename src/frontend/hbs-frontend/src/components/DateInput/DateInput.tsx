import { Box, Button, useDisclosure } from '@chakra-ui/react'
import { useState } from 'react';
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
} from '@chakra-ui/react'
import React from 'react'

import AddIcon from 'components/Icons/Add'
import { DateReservation as DateReservationT } from 'types/DateReservation';
import RoundButton from "components/RoundButton/RoundButton";

import styles from "./DateInput.module.scss"
import DateInputBox from 'components/DateInputBox/DateInputBox'


interface DateContextType {
  startDate: string,
  setStartDate: React.Dispatch<React.SetStateAction<string>>,
  endDate: string,
  setEndDate: React.Dispatch<React.SetStateAction<string>>
}

export const DateContext = React.createContext<DateContextType | undefined>(undefined);


export default function DateInput(props) {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const [ startDate, setStartDate] = useState('');
  const [ endDate, setEndDate] = useState('');

  var data: DateReservationT = { start_date: "", end_date: "" }

  async function put() { 
    data.start_date = startDate;
    data.end_date = endDate;
    await props.putCallback(data)
    
    onClose()
  }

  return (
    <>
      <RoundButton onClick={onOpen}>
          Забронировать
      </RoundButton>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent className={styles.dark_bg}>
          <ModalHeader>Выберете даты въезда и выезда</ModalHeader>
          <ModalCloseButton />
          <ModalBody className={styles.model_body}>
            <Box>
              <DateContext.Provider value={{ startDate, setStartDate, endDate, setEndDate }}>
                <DateInputBox />
              </DateContext.Provider>
            </Box>

            <Button className={styles.ready_btn} onClick={put}>
              <AddIcon className={styles.img_btn} />
            </Button>
          </ModalBody>
          <ModalFooter/>
        </ModalContent>
      </Modal>
    </>
  )
}
