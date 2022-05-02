import { ServerInfo, ContentType } from "./constants";
import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "./firebase";
import type {UserCredential } from "firebase/auth"




const login = async (event: any): Promise<string> => {
    event.preventDefault();

    const { email, password }:
        { email: string, password: string } = event.detail;
    try {
        const userCredential:UserCredential = await signInWithEmailAndPassword(auth, email, password);
        console.log(userCredential.user);
    } catch(error:any) {
        const errorCode = error.code;
        const errorMessage = error.message;
        console.error(`${errorCode}: ${errorMessage}`);
    }
    
    return await Promise.resolve("");
}

const register = async (event: any):Promise<void> => {
    event.preventDefault();
    
    const {displayName, email, password}:
        {displayName:string, email:string, password:string} = event.detail;

    try {
        const userCredential:UserCredential = await createUserWithEmailAndPassword(
            auth, 
            email, 
            password
        );
        console.log(userCredential.user);
    } catch(error:any) {
        const errorCode = error.code;
        const errorMessage = error.message;

        console.error(`${errorCode}: ${errorMessage}`);
    }
}

export {
    login,
    register
}