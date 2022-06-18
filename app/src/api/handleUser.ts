import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from "firebase/auth";
import { auth } from "./firebase";
import type { UserCredential } from "firebase/auth";
import { ServerInfo, ContentType } from "./constants";

const authUrl = `${ServerInfo.serverAddress}/api/user/auth`;
const loginUrl = `${ServerInfo.serverAddress}/api/user/login`;

const login = async (
    email:string, 
    password:string
):Promise<void> => {
    try {
        const userCredential:UserCredential = await signInWithEmailAndPassword(auth, email, password);

        auth.currentUser?.getIdToken(true).then(async (idToken) => {
            const response:Response = await fetch("/api/user/login", {
                method: "POST",
                headers: {
                    ...ContentType.json,
                },
                body: JSON.stringify({
                    'displayName': userCredential.user.displayName,
                    'email': userCredential.user.email,
                    'uid': userCredential.user.uid,
                    'idToken': idToken,
                }),
                
            });

            let json = await response.json();
            console.log(json)
        });
    } catch(error:any) {
        console.error(`${error.code}: ${error.message}`);
    }

}

const register = async (
    displayName:string, 
    email:string, 
    password:string,
):Promise<void> => {
    try {
        const userCredential:UserCredential = await createUserWithEmailAndPassword(
            auth, 
            email, 
            password
        );
    } catch(error:any) {
        console.error(`${error.code}: ${error.message}`);
    }
}

const authUser = async ():Promise<void> => {
    auth.currentUser?.getIdToken(true).then(async () => {
        
    });
}

export {
    login,
    register
}