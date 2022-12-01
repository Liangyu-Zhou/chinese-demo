import React, { FormEvent, useState } from 'react';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import Avatar from '@material-ui/core/Avatar';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { LinearProgress } from '@material-ui/core';
import useSigner from "../hooks/metamaskSigner";
import AddressAvatar from "./AddressAvatar";
import  metamaskImg from "../assets/MetaMask_Fox.svg.png";

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
    },
    form: {
        width: '100%', // IE 11.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
        fontSize: 'small'
    },
}));

const Metamask = () => {
    const classes = useStyles();
    const [inProgress, setInProgress] = useState(false);

    const handleSubmit = (e: FormEvent) => doRequestFunds(e);
    const doRequestFunds = async (e: FormEvent) => {
        e.preventDefault();
    };
    const { address, loading, connectWallet } = useSigner();
    let avatar;
    if (address) {
        avatar = <AddressAvatar address={address}></AddressAvatar>
    } else {
        avatar = <Avatar
            className={classes.avatar}
            alt="1ot found" src={metamaskImg}>
        </Avatar>
    }
    return (
        <>
            {loading && <LinearProgress color="secondary" />}
            <Container component="main" maxWidth="xs">
                <CssBaseline />
                <div className={classes.paper}>
                    {avatar}
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        disabled={loading}
                        className={classes.submit}
                        onClick={connectWallet}
                    >
                        {loading ? "Connecting..." : "Login with Metamask"}
                    </Button>

                </div>
            </Container>
        </>
    );
};

export default Metamask;