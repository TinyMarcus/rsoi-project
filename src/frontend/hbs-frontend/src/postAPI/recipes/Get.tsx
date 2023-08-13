import axios from "axios";
import { Hotel } from "types/Hotel";
import { backUrl } from "..";

interface resp {
    status: number
    content: Hotel
}

const GetRecipe = async function(id: number): Promise<resp> {
    const response = await axios.get(backUrl + `/recipes/${id}`);
    return {
        status: response.status,
        content: response.data as Hotel
    };
}

export default GetRecipe
