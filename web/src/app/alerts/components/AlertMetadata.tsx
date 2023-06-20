import React, { useState, useEffect } from 'react'
import { useQuery, useMutation, gql } from 'urql'
import { Grid, Grow, IconButton, Tooltip, TextField } from '@mui/material'
import {
  ThumbUp,
  ThumbDown,
  ThumbUpOutlined,
  ThumbDownOutlined,
  Info,
} from '@mui/icons-material'
import { DEBOUNCE_DELAY } from '../../config'

const query = gql`
  query AlertMetadataQuery($id: Int!) {
    alert(id: $id) {
      id
      metadata {
        sentiment
        note
      }
    }
  }
`

const mutation = gql`
  mutation UpdateMetadataMutation($input: UpdateAlertMetadataInput!) {
    updateAlertMetadata(input: $input)
  }
`

interface AlertMetadataProps {
  alertID: number
}

export default function AlertMetadata(props: AlertMetadataProps): JSX.Element {
  const { alertID } = props
  const [{ data, fetching, error }] = useQuery({
    query,
    variables: {
      id: alertID,
    },
    requestPolicy: 'cache-first',
  })
  const [note, setNote] = useState(data?.alert?.metadata?.note ?? '')
  const [mutationStatus, commit] = useMutation(mutation)

  // Debounce setting note
  useEffect(() => {
    const t = setTimeout(() => {
      commit({
        input: {
          alertID,
          note,
        },
      })
    }, DEBOUNCE_DELAY)

    return () => clearTimeout(t)
  }, [note])

  useEffect(() => {
    setNote(data?.alert?.metadata?.note ?? '')
  }, [data?.alert?.metadata?.note])

  const thumbUp =
    !fetching && !error && data?.alert?.metadata?.sentiment === 1 ? (
      <ThumbUp />
    ) : (
      <ThumbUpOutlined />
    )

  const isDown = data?.alert?.metadata?.sentiment === -1
  const thumbDown =
    !fetching && !error && isDown ? <ThumbDown /> : <ThumbDownOutlined />

  return (
    <Grid
      container
      spacing={1}
      justifyContent='flex-end'
      alignItems='center'
      sx={{ pr: '8px' }}
    >
      {mutationStatus.error?.message}
      <Grid item>
        <IconButton
          onClick={() => {
            commit({
              input: {
                alertID,
                sentiment: 1,
              },
            })
          }}
          size='large'
        >
          {thumbUp}
        </IconButton>
      </Grid>
      <Grid item>
        <IconButton
          onClick={() => {
            commit({
              input: {
                alertID,
                sentiment: -1,
              },
            })
          }}
          size='large'
        >
          {thumbDown}
        </IconButton>
      </Grid>
      <Grid item>
        <Tooltip title='Was this alert actionable?' placement='top'>
          <Info fontSize='small' sx={{ p: '12px' }} />
        </Tooltip>
      </Grid>
      <Grow in={isDown}>
        <Grid
          item
          xs={12}
          sx={{
            display: !isDown ? 'none' : 'flex',
            justifyContent: 'flex-end',
          }}
        >
          <TextField value={note} onChange={(e) => setNote(e.target.value)} />
        </Grid>
      </Grow>
    </Grid>
  )
}
