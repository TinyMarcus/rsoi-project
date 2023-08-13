import axios from "axios";
import { AllHotelResp, backUrl } from "..";

const GetRecipes = async function(title?: string): Promise<AllHotelResp> {
    const params = { title: title ? title : '' }
    const response = await axios.get(backUrl + `/categories/${title}/recipes`, {params:params});
    return {
        status: response.status,
        content: response.data
    };
}

export default GetRecipes
