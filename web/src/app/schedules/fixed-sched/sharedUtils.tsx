import React, { ReactNode } from 'react'

export interface Value {
  start: string
  end: string
  shifts: Shift[]
}

export interface Shift {
  start: string
  end: string
  user: User
}

export interface User {
  label: string
  value: string
}

// removes bottom margin from content text so form fields
// don't have a bunch of whitespace above them
export const contentText = {
  marginBottom: 0,
}

interface StepContainerProps {
  children: ReactNode
}
export function StepContainer({ children }: StepContainerProps) {
  const bodyStyle = {
    display: 'flex',
    justifyContent: 'center', // horizontal align
    height: 'fit-content',
    width: '100%',
    marginTop: '5%', // slightly lower below dialog title toolbar
  }

  const containerStyle = {
    width: '35%', // ensures form fields don't shrink down too small
  }

  return (
    <div style={bodyStyle}>
      <div style={containerStyle}>{children}</div>
    </div>
  )
}
