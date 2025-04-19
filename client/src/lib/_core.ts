import { Preferences } from '@capacitor/preferences'
import { user,error,API_URL } from '$lib'


export async function resetUser(){
    user.set({})
    await Preferences.clear()
}

export async function fetchAPI(
    path: string, 
    method: string,
    data?:any
):Promise<any>{
    let result  
    const headers = new Headers()
    const token = await Preferences.get({key: "user"})
    if(token["value"] !== null){
        headers.append("Authorization", "Bearer "+token["value"])
    }
    const options:any = {
        method: method,
        redirect: "follow",
        headers: headers
    }
    if(method === "POST" && data){
        options.body = data
    }
    try{
        const response = await fetch(API_URL+path, options)
        result = await response.json()
        if(!response.ok){
            if(response.status === 401){
                return {
                    "error": result.error
                }
            }
            throw result.error
        }
    } catch(err:any){
        error.set(err)
        return {
            "error": err
        }
    }
    return result
}