import React, { useState, useEffect } from 'react'
import { Box } from '@mui/system'
import { DateTime } from 'luxon'
import { DebugMessage } from '../../../schema'
import OutgoingLogCard from './OutgoingLogCard'
import { useFuse } from './hooks'
import { useURLParam } from '../../actions'
import { Typography, Button } from '@mui/material'

const INITIAL_LIMIT = 1
const LOAD_AMOUNT = 8

interface Props {
  debugMessages?: DebugMessage[]
  onSelect: (debugMessage: DebugMessage) => void
}

export default function OutgoingLogsList(props: Props): JSX.Element {
  const { debugMessages = [], onSelect } = props

  const [searchTerm] = useURLParam('search', '')
  const [start] = useURLParam('start', '')
  const [end] = useURLParam('end', '')

  const [limit, setLimit] = useState(INITIAL_LIMIT)

  const { setSearch, results } = useFuse<DebugMessage>({
    data: debugMessages,
    keys: ['status'], // todo: add more keys, phone number/service/user name/etc
    options: { shouldSort: false },
    customOptions: { showResultsWhenNoSearchTerm: true },
  })

  console.log(searchTerm, debugMessages, results)

  let filteredResults = results.slice() // copy results array
  filteredResults = filteredResults.filter((result) => {
    if (!start && !end) return true

    const startDT = DateTime.fromISO(start)
    let endDT = DateTime.fromISO(end)
    const createdAtDT = DateTime.fromISO(result.item.createdAt)

    if (start && !end) {
      endDT = DateTime.now()
    }

    if (createdAtDT > startDT && createdAtDT < endDT) {
      return true
    }

    return false
  })

  useEffect(() => {
    setLimit(INITIAL_LIMIT)
  }, [searchTerm, start, end])

  useEffect(() => {
    setSearch(searchTerm)
  }, [searchTerm])

  // what appends stuff to results
  function onNext(): void {
    console.log('onNext called', limit * LOAD_AMOUNT)
    setLimit(limit + 1)
    console.log('onNext called (after)', limit * LOAD_AMOUNT)
  }

  // function hasMore(): boolean {
  //   if (filteredResults.length > limit * LOAD_AMOUNT) {
  //     return true
  //   }

  //   return false
  // }

  return (
    <Box
      display='flex'
      flexDirection='column'
      alignItems='stretch'
      width='full'
    >
      {/* TODO: change card's outline color in list when selected */}
      {filteredResults
        .slice(0, limit * LOAD_AMOUNT)
        .map(({ item: debugMessage }, index) => (
          <OutgoingLogCard
            key={debugMessage.id}
            index={index}
            debugMessage={debugMessage}
            onClick={() => onSelect(debugMessage)}
          />
        ))}
      {limit * LOAD_AMOUNT < filteredResults.length ? (
        // load more
        <Button
          variant='contained'
          color='primary'
          sx={{ marginTop: '1rem', marginBottom: '1rem' }}
          onClick={onNext}
        >
          Load more...
        </Button>
      ) : (
        // done loading
        <Typography
          color='textSecondary'
          variant='body2'
          sx={{
            display: 'flex',
            justifyContent: 'center',
            padding: '0.25em 0 0.25em 0',
          }}
        >
          Displaying all results.
        </Typography>
      )}
    </Box>
  )
}
