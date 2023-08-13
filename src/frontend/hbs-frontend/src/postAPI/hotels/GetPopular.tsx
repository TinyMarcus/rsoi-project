import axiosBackend, { PopularHotelsResp } from ".."

const GetPopular = async function(): Promise<PopularHotelsResp> {
    const response = await axiosBackend
        .get("/hotels/popular");

    return {
        status: response.status,
        content: response.data
    };
}

export default GetPopular
