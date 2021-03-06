import { Model } from "./Model";

export class User extends Model {
    private email: string;
    private username: string;

    constructor({ id, email, username }:
        { id: number, email: string, username: string }) {

        super(id);
        this.username = username;
        this.email = email;
    }

    public getEmail(): string {
        return this.email;
    }

    public getUsername(): string {
        return this.username;
    }
}