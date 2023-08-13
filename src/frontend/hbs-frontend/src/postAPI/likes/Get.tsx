import axiosBackend from "..";

interface resp {
    status: number
    content: string
}

const GetImageUrl = async function(hotel_uid: string): Promise<resp> {
    const response = await axiosBackend.get(`/hotels/${hotel_uid}/image`);
    return {
        status: response.status,
        content: response.data as string
    };
}
export default GetImageUrl;
