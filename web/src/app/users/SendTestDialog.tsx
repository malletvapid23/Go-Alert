import React, { useEffect, useState, MouseEvent } from 'react'
import { useQuery, useMutation, gql } from '@apollo/client'

import Spinner from '../loading/components/Spinner'

import {
  Button,
  Dialog,
  DialogActions,
  DialogTitle,
  DialogContent,
  DialogContentText,
} from '@mui/material'
import toTitleCase from '../util/toTitleCase'
import DialogContentError from '../dialogs/components/DialogContentError'
import { DateTime } from 'luxon'
import { ContactMethodType } from '../../schema'

const query = gql`
  query ($id: ID!) {
    userContactMethod(id: $id) {
      id
      type
      formattedValue
      lastTestVerifyAt
      lastTestMessageState {
        details
        status
        formattedSrcValue
      }
    }
  }
`

const mutation = gql`
  mutation ($id: ID!) {
    testContactMethod(id: $id)
  }
`

export default function SendTestDialog(
  props: SendTestDialogProps,
): JSX.Element {
  const { title = 'Test Delivery Status', onClose, messageID } = props

  const [now] = useState(DateTime.local())

  const [sendTest, sendTestStatus] = useMutation(mutation, {
    variables: {
      id: messageID,
    },
  })

  const { data, loading, error } = useQuery(query, {
    variables: {
      id: messageID,
    },
    fetchPolicy: 'network-only',
  })

  const status = data?.userContactMethod?.lastTestMessageState?.status ?? ''
  const cmDestValue = data?.userContactMethod?.formattedValue ?? ''
  const cmType: ContactMethodType = data?.userContactMethod?.type ?? ''
  const lastTestVerifyAt = data?.userContactMethod?.lastTestVerifyAt ?? ''
  const timeSinceLastVerified = now.diff(DateTime.fromISO(lastTestVerifyAt))
  const fromValue =
    data?.userContactMethod?.lastTestMessageState?.formattedSrcValue ?? ''
  const errorMessage = error?.message ?? ''

  useEffect(() => {
    if (loading || error || sendTestStatus.called) {
      return
    }
    if (
      data?.userContactMethod?.lastTestMessageState == null ||
      !(timeSinceLastVerified.as('seconds') < 60)
    ) {
      sendTest()
    }
  }, [lastTestVerifyAt, loading])

  let details
  if (sendTestStatus.called && lastTestVerifyAt > now.toISO()) {
    details = data?.userContactMethod?.lastTestMessageState?.details ?? ''
  }

  const getTestStatusColor = (status: string): string => {
    switch (status) {
      case 'OK':
        return 'success'
      case 'ERROR':
        return 'error'
      default:
        return 'warning'
    }
  }

  if (loading || sendTestStatus.loading) return <Spinner text='Loading...' />

  const msg = (): string => {
    switch (cmType) {
      case 'SMS':
      case 'VOICE':
        return `${
          cmType === 'SMS' ? 'SMS message' : 'voice call'
        } to ${cmDestValue}`
      case 'EMAIL':
        return `email to ${cmDestValue}`
      default:
        return `to ${cmDestValue}`
    }
  }

  return (
    <Dialog open onClose={onClose}>
      <DialogTitle>{title}</DialogTitle>

      <DialogContent>
        <DialogContentText>
          GoAlert is sending a test {msg()}.
        </DialogContentText>
        {fromValue && (
          <DialogContentText>
            The test message was sent from {fromValue}.
          </DialogContentText>
        )}
        {details && (
          <DialogContentText color={getTestStatusColor(status)}>
            {toTitleCase(details)}
          </DialogContentText>
        )}
        {!details && (
          <DialogContentText color='error'>
            Couldn't send a message yet, please try again after about a minute.
          </DialogContentText>
        )}
      </DialogContent>

      {errorMessage && <DialogContentError error={errorMessage} />}

      <DialogActions>
        <Button color='primary' variant='contained' onClick={onClose}>
          Done
        </Button>
      </DialogActions>
    </Dialog>
  )
}

interface SendTestDialogProps {
  messageID: string
  onClose: (event: MouseEvent) => void
  disclaimer?: string
  title?: string
  subtitle?: string
}
