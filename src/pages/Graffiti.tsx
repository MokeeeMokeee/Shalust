import React from 'react'
import styled from 'styled-components'

/**
 * components
 */
import Items from '../components/Items'
import Navbar from '../components/Navbar'

/**
 * testData
 */
import { GraffitiData } from '../data/Data'

const Container = styled.div`
  min-height: 80vh; /* 後で消す */
  padding: 2rem 1rem;
  background-color: #e6ecf0;
`
const Content = styled.div`
  padding: 2rem 1rem;
  background-color: #fff;
`

const Graffiti: React.FC = () => {
  return (
    <Container>
      <Content>
        <Navbar selectNumber={4} />
        <Items datas={GraffitiData} isInfo={true} />
      </Content>
    </Container>
  )
}

export default Graffiti
