// @flow
import { GetServerSideProps, NextPage } from "next";
import * as React from "react";
type Props = {};
type State = {};

const Pagina1Page: NextPage = () => {
  return (
    <div>
      <h1>Ola Mundo</h1>
    </div>
  );
};

const getServerSideProps: GetServerSideProps = async (context) => {
  return {
    props: {
      name: "João Dematé",
    },
  };
};

export default Pagina1Page;
