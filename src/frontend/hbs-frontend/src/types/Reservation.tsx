import { HotelInfo } from "./HotelInfo";
import { PaymentInfo } from "./PaymentInfo";

export interface Reservation {
	reservation_uid:            string,
    hotel:                      HotelInfo,
    start_date:                 string,
    end_date:                   string,
    status:                     string,
    payment:                    PaymentInfo
}