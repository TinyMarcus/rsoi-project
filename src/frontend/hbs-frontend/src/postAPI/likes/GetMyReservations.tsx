import axiosBackend from ".."
import { AllHotelResp } from "..";

const GetMyReservations = async function(): Promise<AllHotelResp> {
    const response = await axiosBackend
        .get(`/reservations`);

    return {
        status: response.status,
        content: response.data
    };
}

export default GetMyReservations